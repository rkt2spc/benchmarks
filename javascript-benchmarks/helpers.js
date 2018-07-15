// Dependencies
const uuid = require('uuid');

// Sample UUIDs
const sample = [];

// Get Sample UUIDs
exports.getSampleUUIDs = (size) => {
  if (size < sample.length) return sample.slice(0, size);
  while (sample.length < size) sample.push(uuid.v4());
  return sample;
};
