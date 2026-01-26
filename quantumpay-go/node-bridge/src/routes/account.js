const express = require('express');
const router = express.Router();
const client = require('../client');

router.get('/:address', (req, res) => {
  client.GetBalance(
    { address: req.params.address },
    (err, r) => {
      if (err) return res.status(500).json({ error: err.message });
      res.json(r);
    }
  );
});

module.exports = router;
