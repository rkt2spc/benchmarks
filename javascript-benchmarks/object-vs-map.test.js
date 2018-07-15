/**
 * @jest-environment node
 */

// Dependencies
const Benchmark = require('benchmark');
const uuid = require('uuid');
const helpers = require('./helpers');

// Benchmark building vanilla Object vs ES6 Map
test('Benchmark building vanilla Object vs ES6 Map', (done) => {
  // Benchmark.prototype.setup = function setup() {
  //   this.keys = Array(this.count).fill().map(() => uuid.v4());
  // };

  const suite = new Benchmark.Suite();
  suite.add('Build Object', function buildObject() {
    const obj = {};
    this.keys.forEach((k) => {
      obj[k] = true;
    });
  }, {
    setup: function setup() { this.keys = helpers.getSampleUUIDs(this.count); },
  });

  suite.add('Build Map', function buildMap() {
    const map = new Map();
    this.keys.forEach((k) => {
      map.set(k, true);
    });
  }, {
    setup: function setup() { this.keys = helpers.getSampleUUIDs(this.count); },
  });

  suite.on('cycle', (event) => console.log('Cycle:', String(event.target)));
  suite.on('complete', () => {
    console.log('Xong');
    done();
  });

  suite.run();
});
