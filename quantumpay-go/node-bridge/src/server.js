const express = require('express');
const health = require('./routes/health');
const chain = require('./routes/chain');
const account = require('./routes/account');

const app = express();
app.use(express.json());

app.use('/health', health);
app.use('/chain', chain);
app.use('/account', account);

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Node Bridge REST listening on :${PORT}`);
});
