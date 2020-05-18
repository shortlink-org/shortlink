import { Fragment, useState } from 'react'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import Snackbar from '@material-ui/core/Snackbar'
import Alert from '@material-ui/lab/Alert'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'
import IconButton from '@material-ui/core/IconButton'
import FileCopyIcon from '@material-ui/icons/FileCopy'
import Grid from '@material-ui/core/Grid'
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    overflow: 'hidden',
    padding: theme.spacing(0, 3),
  },
  form: {
    display: 'grid',
  },
  paper: {
    maxWidth: 400,
    margin: `${theme.spacing(1)}px auto`,
    padding: theme.spacing(2),
  },
}));

export default function AddURL() {
  const [open, setOpen] = React.useState(false)
  const classes = useStyles()

  const [url, setURL] = useState({
    url: '',
  });

  const [response, setResponse] = useState({
    type: '',
    message: '',
    hash: ''
  })

  const handleChange = e =>
    setURL({ ...url, [e.target.name]: e.target.value })

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      const res = await fetch(`${process.env.API_URL_HTTP}/api`, {
        method: 'POST',
        body: JSON.stringify(url),
        headers: { 'Content-Type': 'application/json' },
      })
      const json = await res.json()

      if (res.status === 201) {
        setResponse({
          type: 'success',
          message: 'Success add your link.',
          hash: json.hash,
        });
      } else {
        setResponse({
          type: 'error',
          message: json.error,
          hash: '',
        });
      }

      setOpen(true)
    } catch (e) {
      console.error('An error occurred', e);
      setResponse({
        type: 'error',
        message: 'An error occured while submitting the form'
      });
      setOpen(true)
    }
  };

  return (
    <Grid container direction="column" justify="space-around" alignItems="center" className={classes.root}>
      <Paper className={classes.paper}>
        <form autoComplete="off" onSubmit={handleSubmit} className={classes.form}>
          <TextField label="Your URL" name="url" onChange={handleChange} />
          <TextField label="Describe" name="describe" onChange={handleChange} />
          <Button variant="contained" color="primary" type={"submit"}>
            Add
          </Button>
        </form>
      </Paper>

      {
        response.type !== "" && (
          <Paper elevation={3} className={classes.paper}>
            <Typography variant="p" component="p">
              Your link: {response.hash}

              <IconButton aria-label="copy" color="secondary">
                <FileCopyIcon />
              </IconButton>
            </Typography>
          </Paper>
        )
      }

      <Snackbar
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'left',
        }}
        open={open}
        autoHideDuration={6000}
        onClose={handleClose}
      >
        <Alert onClose={handleClose} severity={response.type}>
          { response.message }
        </Alert>
      </Snackbar>
    </Grid>
  );
}
