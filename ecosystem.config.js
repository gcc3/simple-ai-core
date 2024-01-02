module.exports = {
  apps : [{
    name: 'node-test',
    script: 'go',
    args: 'run serve.go',
    // watch: true,   // restart on file changes
    env: {
      PORT: 3030
    }
  }]
};
