const express = require('express');
const router = express.Router();
const client = require('../client');

router.get('/', (_, res) => {
  client.Health({}, (err, r) => {
    if (err) return res.status(500).json({ error: err.message });
    res.json(r);
  });
});

module.exports = router;
