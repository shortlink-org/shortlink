const { exec } = require('child_process');
const util = require('util');
const yargs = require('yargs/yargs');
const { hideBin } = require('yargs/helpers');
const fs = require('fs');
const os = require('os');

const execPromisified = util.promisify(exec);

async function main() {
  const argv = yargs(hideBin(process.argv)).argv;
  const namespace = argv.namespace;
  const secretName = argv.secret;
  const key = argv.key;
  const envKey = argv.envKey;

  if (!namespace || !secretName || !key || !envKey) {
    console.log("Please provide a namespace, secret name, a key, and an envKey as command line arguments.");
    process.exit(1);
  }

  try {
    const { stdout } = await execPromisified(`kubectl -n ${namespace} get secret ${secretName} -o json`);
    const secret = JSON.parse(stdout);
    const keyValueBase64 = secret.data[key];
    const keyValue = Buffer.from(keyValueBase64, 'base64').toString();
    const keyValueDecoded = decodeURIComponent(keyValue);

    let envConfig = fs.readFileSync('.env', 'utf8').split(os.EOL);
    let newConfig = envConfig.map(line => {
      if (line.startsWith(envKey)) {
        return `${envKey}=${keyValueDecoded}`
      }
      return line;
    });

    if (!newConfig.includes(`${envKey}=${keyValueDecoded}`)) {
      newConfig.push(`${envKey}=${keyValueDecoded}`);
    }

    fs.writeFileSync('.env', newConfig.join(os.EOL));
    console.log(`Written to .env: ${envKey}=${keyValueDecoded}`);
  } catch (error) {
    console.error(`Error: ${error}`);
  }
}

main();
