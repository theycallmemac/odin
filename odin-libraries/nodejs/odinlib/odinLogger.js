const http = require('http');

class OdinLogger {
  async log(type, desc, value, id, timestamp) {
    const data = type + ',' + desc + ',' + value + ',' + id + ',' + timestamp;

    const options = {
      host: 'localhost',
      port: 3939,
      path: '/stats/add',
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Content-Length': data.length,
      },
    };

    const request = http.request(options, (response) => {
      response.on('data', (d) => {
        process.stdout.write(d);
      });
    });

    request.on('error', (error) => {
      return false;
    });

    request.write(data);
    request.end();
    return true;
  }
}

module.exports.OdinLogger = OdinLogger;

