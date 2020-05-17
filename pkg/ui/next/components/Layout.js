import React from 'react';
import Container from '@material-ui/core/Container';
import Header from './Header';
import Copyright from './Copyright';

const Layout = props => (
  <Container maxWidth="sm">
    <Header />
    {props.content}
    <Copyright />
  </Container>
);

export default Layout;
