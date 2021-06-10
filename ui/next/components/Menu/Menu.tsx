import IconButton from '@material-ui/core/IconButton'
import clsx from 'clsx'
import Drawer from '@material-ui/core/Drawer'
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft'
import Divider from '@material-ui/core/Divider'
import List from '@material-ui/core/List'
import React from 'react'
import { mainListItems, secondaryListItems, adminListItems } from './listItems'
import useStyles from './style'
import {useSelector} from "react-redux";

const Menu = ({ open, setOpen }) => {
  const classes = useStyles()

  // @ts-ignore
  const session = useSelector(state => state.session)

  const handleDrawerClose = () => {
    setOpen(false)
  }

  return (
    <Drawer
      variant="permanent"
      classes={{
        paper: clsx(classes.drawerPaper, !open && classes.drawerPaperClose),
        root: classes.root,
      }}
      open={open}
    >
      <div className={classes.toolbarIcon}>
        <IconButton onClick={handleDrawerClose}>
          <ChevronLeftIcon />
        </IconButton>
      </div>
      <Divider />

      <List>{mainListItems}</List>
      <Divider />

      <List>{secondaryListItems}</List>
      <Divider />

      <List>{adminListItems}</List>
    </Drawer>
  )
}

export default Menu
