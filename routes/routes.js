const controller = require('../controller/controller');

module.exports = (app) => {
    app.get('/api/v1/listCountries', controller.listCountries);
    app.get('/api/v1/indStats', controller.indStats);
    app.get('/api/v1/indDistrict', controller.indDistrict);
    app.get('/api/v1/dataJson', controller.dataJson);
    app.get('/', controller.pingHealth);
}