const Menu = ({ open, setOpen }) => {
  const session = useSelector((state) => state.session)

  if (!session.kratos.active) {
    return null
  }

  return (<></>)
}

export default Menu
