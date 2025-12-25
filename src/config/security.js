module.exports = {
  BACKOFF: {
    BASE_DELAY_MS: 2000,      // 2 detik
    MAX_DELAY_MS: 60000,      // 60 detik max
    STRIKE_RESET_MS: 300000,  // reset setelah 5 menit
    MAX_STRIKES: 10
  }
};

