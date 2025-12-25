'use strict';

const fs = require('fs');
const path = require('path');

let wasmInstance = null;

async function loadWasm() {
  if (wasmInstance) return wasmInstance;

  const wasmPath = path.join(
    __dirname,
    '../wasm/liboqs_dilithium.wasm'
  );

  const buffer = fs.readFileSync(wasmPath);
  const module = await WebAssembly.compile(buffer);
  const instance = await WebAssembly.instantiate(module, {});

  wasmInstance = instance.exports;
  return wasmInstance;
}

/**
 * Expected WASM exports (liboqs):
 * - dilithium_keypair()
 * - dilithium_sign(msgPtr, msgLen)
 * - dilithium_verify(msgPtr, msgLen, sigPtr)
 */

async function generateKeypair() {
  const wasm = await loadWasm();
  return wasm.dilithium_keypair();
}

async function sign(message) {
  const wasm = await loadWasm();
  return wasm.dilithium_sign(message);
}

async function verify(message, signature) {
  const wasm = await loadWasm();
  return wasm.dilithium_verify(message, signature);
}

module.exports = {
  generateKeypair,
  sign,
  verify
};

