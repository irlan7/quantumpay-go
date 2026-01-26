const express = require('express');
const router = express.Router();
const client = require('../client');

router.get('/height', (_, res) => {
  client.GetHeight({}, (err, r) => {
    if (err) return res.status(500).json({ error: err.message });
    res.json(r);
  });
});

router.get('/block/:height', (req, res) => {
  client.GetBlock(
    { height: Number(req.params.height) },
    (err, r) => {
      if (err) return res.status(500).json({ error: err.message });
      res.json(r);
    }
  );
});

module.exports = router;
