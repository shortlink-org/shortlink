package v1

// LeaderPromotion promotes the current node to leader.
func (r *Raft) LeaderPromotion() {
	r.status = RaftStatus_RAFT_STATUS_LEADER
}
