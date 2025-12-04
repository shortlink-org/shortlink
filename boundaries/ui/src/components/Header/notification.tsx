import NotificationsIcon from '@mui/icons-material/Notifications'
import Badge from '@mui/material/Badge'
import Fade from '@mui/material/Fade'
import IconButton from '@mui/material/IconButton'
import Menu from '@mui/material/Menu'
import * as React from 'react'

export default function Notification(): React.JSX.Element {
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null)
  const open = Boolean(anchorEl)
  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget)
  }
  const handleClose = () => {
    setAnchorEl(null)
  }

  return (
    <>
      <IconButton
        color="inherit"
        aria-controls="notification-menu"
        aria-haspopup="true"
        aria-expanded={open ? 'true' : undefined}
        onClick={handleClick}
        size="large"
        className="hover:bg-white/10 active:bg-white/20 transition-all duration-200 rounded-lg"
        sx={{
          '&:hover': {
            transform: 'scale(1.05)',
          },
          '&:active': {
            transform: 'scale(0.95)',
          },
        }}
      >
        <Badge 
          badgeContent={4} 
          color="error"
          sx={{
            '& .MuiBadge-badge': {
              backgroundColor: '#ef4444',
              color: 'white',
              fontSize: '0.75rem',
              fontWeight: 600,
            }
          }}
        >
          <NotificationsIcon className="text-white" />
        </Badge>
      </IconButton>

      <Menu
        id="notification-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        onClose={handleClose}
        TransitionComponent={Fade}
        sx={{ 
          padding: 0,
          '& .MuiPaper-root': {
            backgroundColor: 'transparent',
            boxShadow: 'none',
            marginTop: '8px',
          }
        }}
      >
        {/* Dropdown menu */}
        <div className="right-0 z-20 overflow-hidden bg-white dark:bg-gray-800 rounded-xl shadow-xl ring-1 ring-black/5 dark:ring-white/10 w-80 max-h-96">
          <div className="bg-gradient-to-r from-indigo-50 to-purple-50 dark:from-gray-700 dark:to-gray-600 px-4 py-3 border-b border-gray-100 dark:border-gray-600">
            <h3 className="text-sm font-semibold text-gray-900 dark:text-white">Notifications</h3>
            <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">You have 4 new notifications</p>
          </div>
          
          <div className="py-2 max-h-64 overflow-y-auto">
            <a className="flex items-start gap-3 px-4 py-3 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-200 border-b border-gray-100 dark:border-gray-700 last:border-b-0">
              <img
                className="object-cover w-8 h-8 rounded-full ring-2 ring-gray-200 dark:ring-gray-600"
                src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=334&q=80"
                alt="Sara Salah"
              />
              <div className="flex-1 min-w-0">
                <p className="text-sm text-gray-900 dark:text-white leading-relaxed">
                  <span className="font-semibold">Sara Salah</span> replied on the{' '}
                  <span className="font-semibold text-indigo-600 dark:text-indigo-400">Upload Image</span> article.
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">2 minutes ago</p>
              </div>
            </a>
            
            <a className="flex items-start gap-3 px-4 py-3 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-200 border-b border-gray-100 dark:border-gray-700 last:border-b-0">
              <img
                className="object-cover w-8 h-8 rounded-full ring-2 ring-gray-200 dark:ring-gray-600"
                src="https://images.unsplash.com/photo-1531427186611-ecfd6d936c79?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=634&q=80"
                alt="Slick Net"
              />
              <div className="flex-1 min-w-0">
                <p className="text-sm text-gray-900 dark:text-white leading-relaxed">
                  <span className="font-semibold">Slick Net</span> started following you.
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">45 minutes ago</p>
              </div>
            </a>
            
            <a className="flex items-start gap-3 px-4 py-3 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-200 border-b border-gray-100 dark:border-gray-700 last:border-b-0">
              <img
                className="object-cover w-8 h-8 rounded-full ring-2 ring-gray-200 dark:ring-gray-600"
                src="https://images.unsplash.com/photo-1450297350677-623de575f31c?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=334&q=80"
                alt="Jane Doe"
              />
              <div className="flex-1 min-w-0">
                <p className="text-sm text-gray-900 dark:text-white leading-relaxed">
                  <span className="font-semibold">Jane Doe</span> liked your reply on{' '}
                  <span className="font-semibold text-indigo-600 dark:text-indigo-400">Test with TDD</span> article.
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">1 hour ago</p>
              </div>
            </a>
            
            <a className="flex items-start gap-3 px-4 py-3 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-200">
              <img
                className="object-cover w-8 h-8 rounded-full ring-2 ring-gray-200 dark:ring-gray-600"
                src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=398&q=80"
                alt="Abigail Bennett"
              />
              <div className="flex-1 min-w-0">
                <p className="text-sm text-gray-900 dark:text-white leading-relaxed">
                  <span className="font-semibold">Abigail Bennett</span> started following you.
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">3 hours ago</p>
              </div>
            </a>
          </div>
          
          <a className="block py-3 font-semibold text-center text-white bg-purple-700 dark:bg-indigo-500 hover:bg-purple-800 dark:hover:bg-indigo-600 transition-all duration-200 cursor-pointer">
            See all notifications
          </a>
        </div>
      </Menu>
    </>
  )
}
