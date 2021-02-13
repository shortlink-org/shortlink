import Button from '@material-ui/core/Button';
import { Layout } from '../components';

const aboutPageContent = <div>
  <p>This is the about page</p>
</div>;

export default function About() {
  return <Layout content={aboutPageContent} />;
}
