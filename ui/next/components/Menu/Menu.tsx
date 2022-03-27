// @ts-nocheck

import IconButton from '@mui/material/IconButton'
import clsx from 'clsx'
import Drawer from '@mui/material/Drawer'
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft'
import Divider from '@mui/material/Divider'
import List from '@mui/material/List'
import React from 'react'
import { useSelector } from 'react-redux'
import { mainListItems, secondaryListItems, adminListItems } from './listItems'
import useStyles from './style'

const Menu = ({ open, setOpen }) => {
  const classes = useStyles()

  // @ts-ignore
  const session = useSelector((state) => state.session)

  const handleDrawerClose = () => {
    setOpen(false)
  }

  if (!session.kratos.active) {
    return null
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
        <IconButton onClick={handleDrawerClose} size="large">
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
