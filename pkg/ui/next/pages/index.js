import Layout from '../components/MyLayout.js';
import AddURL from '../components/AddURL.js';

export default function Index() {
  return <Layout content={<AddURL />} />;
}
