import Link from 'next/link';
import Typography from '@material-ui/core/Typography';

const linkStyle = {
  marginRight: 15
};

const Header = () => (
  <div>
    <Typography variant="h4" component="h1" gutterBottom>
      Shortlink
    </Typography>

    <Link href="/">
      <a style={linkStyle}>Home</a>
    </Link>

    <Link href="/about">
      <a style={linkStyle}>About</a>
    </Link>
  </div>
);

export default Header;
