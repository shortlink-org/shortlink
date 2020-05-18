import Link from 'next/link'
import { useRouter } from 'next/router'
import AppBar from '@material-ui/core/AppBar'
import Tab from '@material-ui/core/Tab'
import Tabs from '@material-ui/core/Tabs'
import Toolbar from '@material-ui/core/Toolbar'
import IconButton from '@material-ui/core/IconButton'
import Typography from '@material-ui/core/Typography'
import MenuIcon from '@material-ui/icons/Menu'
import SearchForm from './SearchForm'
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
  title: {
    flexGrow: 1,
  },
}));

const Header = () => {
  const router = useRouter()
  const classes = useStyles()

  return (
    <AppBar position="static" value={router.route}>
      <Toolbar>
        <IconButton edge="start" color="inherit" aria-label="menu">
          <MenuIcon />
        </IconButton>

        <Typography variant="h6" className={classes.title}>
          Shortlink
        </Typography>

        <Tabs>
          <Link href="/">
            <Tab label="Home" value="/" />
          </Link>

          <Link href="/list">
            <Tab label="List" value="/" />
          </Link>

          <Link href="/about">
            <Tab label="About" value="/about" />
          </Link>
        </Tabs>

        <SearchForm />
      </Toolbar>
    </AppBar>
  )
};

export default Header;
