#include "_cgo_export.h"

#define QUEUE_DEPTH 16
#define INITIAL_NUM_BLOCKS 7

struct io_uring ring;

struct file_info {
  __u8 opcode; /* type of operation for this sqe */
  off_t file_sz;
  int file_fd;
  struct iovec iovecs[];
};

int pop_request() {
  struct io_uring_cqe *cqe;
  // get from queue
  int ret = io_uring_peek_cqe(&ring, &cqe);
  if (ret < 0) {
    fprintf(stderr, "bad ret.\n");
    return ret;
  }
  if (cqe->res < 0) {
    fprintf(stderr, "bad res.\n");
    return cqe->res;
  }

  struct file_info *fi = io_uring_cqe_get_data(cqe);
  if (fi->opcode == IORING_OP_READV) {
    int total_blocks = INITIAL_NUM_BLOCKS;
    off_t initial_block_size = fi->file_sz / total_blocks;

    if (initial_block_size == 0) {
        total_blocks = 0;
    }

    if (fi->file_sz - initial_block_size*total_blocks > 0) {
      total_blocks++;
    }
    // call read_callback to Go
    read_callback(fi->iovecs, total_blocks, fi->file_fd);

    for (int i=0; i < total_blocks; i++) {
      free(fi->iovecs[i].iov_base);
    }
  } else if (fi->opcode == IORING_OP_WRITEV) {
    // call write_callback to Go
    write_callback(cqe->res, fi->file_fd);
  }
  free(fi);

  // mark as done.
  io_uring_cqe_seen(&ring, cqe);
  return 0;
}

int push_read_request(int file_fd, off_t file_sz) {
  // aiming for 8 blocks. (https://www.oreilly.com/library/view/linux-system-programming/9781449341527/ch04.html)
  int total_blocks = INITIAL_NUM_BLOCKS;
  off_t last_block_size;
  off_t initial_block_size;

  last_block_size = initial_block_size = file_sz / total_blocks;

  if (initial_block_size == 0) {
    total_blocks = 0;
  }

  if (file_sz - initial_block_size*total_blocks > 0) {
    last_block_size = file_sz - initial_block_size*total_blocks;
    total_blocks++;
  }

  struct file_info *fi = malloc(sizeof(*fi) + (sizeof(struct iovec) * total_blocks));
  // populate iovecs
  for (int i=0; i < total_blocks; i++) {
    off_t current_block_size = initial_block_size;
    if (i == total_blocks-1) {
      // last block, need to change the block size
      current_block_size = last_block_size;
    }
    fi->iovecs[i].iov_len = current_block_size;
    void *buf;
    if (posix_memalign(&buf, 1024, current_block_size)) {
      perror("posix_memalign");
      return -1;
    }
    fi->iovecs[i].iov_base = buf;
  }

  fi->file_sz = file_sz;
  fi->file_fd = file_fd;
  fi->opcode = IORING_OP_READV;

  // set the queue
  struct io_uring_sqe *sqe = io_uring_get_sqe(&ring);
  io_uring_prep_readv(sqe, file_fd, fi->iovecs, total_blocks, 0);
  io_uring_sqe_set_data(sqe, fi);
  return 0;
}

int push_write_request(int file_fd, void *data, off_t file_sz) {
  struct file_info *fi = malloc(sizeof(*fi) + (sizeof(struct iovec) * 1));

  fi->iovecs[0].iov_base = data;
  fi->iovecs[0].iov_len = file_sz;
  fi->file_sz = file_sz;
  fi->file_fd = file_fd;
  fi->opcode = IORING_OP_WRITEV;

  // set the queue
  struct io_uring_sqe *sqe = io_uring_get_sqe(&ring);
  io_uring_prep_writev(sqe, file_fd, fi->iovecs, 1, 0);
  io_uring_sqe_set_data(sqe, fi);
  return 0;
}

int queue_submit(int num) {
  return io_uring_submit_and_wait(&ring, num);
}

int queue_init() {
  return io_uring_queue_init(QUEUE_DEPTH, &ring, 0);
}

void queue_exit() {
  io_uring_queue_exit(&ring);
}
