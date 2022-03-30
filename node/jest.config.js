const {defaults} = require('jest-config');
module.exports = {
    preset: 'ts-jest',
    testEnvironment: 'node',
    setupFiles: ['dotenv/config'],
    moduleFileExtensions: [...defaults.moduleFileExtensions, 'd.ts']
};