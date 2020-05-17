import Layout from '../components/Layout.js';
import AddURL from '../components/AddURL.js';

export default function Index() {
  return <Layout content={<AddURL />} />;
}
