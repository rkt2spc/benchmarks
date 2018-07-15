// Dependencies
const Benchmark = require('benchmark');
const uuid = require('uuid');
const helpers = require('./helpers');

const suite = new Benchmark.Suite();
suite.add('Build Object', function buildObject() {
  const obj = {};
  this.keys.forEach((k) => {
    obj[k] = true;
  });
}, {
  setup: function() {
    console.log(this);
    this.keys = helpers.getSampleUUIDs(this.count);
  },
});

// suite.add('Build Map', function buildMap() {
//   const map = new Map();
//   this.keys.forEach((k) => {
//     map.set(k, true);
//   });
// }, { setup });

suite.on('cycle', function(event) {
  console.log(String(event.target));
});
// suite.on('complete', () => {
//   console.log('Xong');
// });

suite.run();
