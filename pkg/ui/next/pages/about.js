import Button from '@material-ui/core/Button';
import Layout from '../components/Layout.js';
import sentry from "../utils/sentry";

const { Sentry, captureException } = sentry()

const aboutPageContent = <div>
  <p>This is the about page</p>
  <Button variant="contained" color="secondary" onClick={() => captureException("123")}>
    Try error
  </Button>
</div>;

export default function About() {
  return <Layout content={aboutPageContent} />;
}
