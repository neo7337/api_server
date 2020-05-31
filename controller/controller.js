const helper = require('../helper/helper');
const configjs = require('../config/config');
const config = configjs.config();

exports.dataJson = (req, res) => {
    helper.makeRestCall(config.indHost, config.historicalINDPath).then((respData) => {
        return respData;
    }).then((resp) => {
        let jsonData = JSON.parse(resp.body);
        //console.log(jsonData);
        return res.json(jsonData);
    }).catch((err) => {
        return res.json(err);
    })
}

exports.indDistrict = (req, res) => {
    helper.makeRestCall(config.indHost, config.districtIndPath).then((respData) => {
        return respData;
    }).then((resp) => {
        let jsonData = JSON.parse(resp.body);
        //console.log(jsonData);
        return res.json(jsonData);
    }).catch((err) => {
        return res.json(err);
    })
}

exports.indStats = (req, res) => {
    helper.makeRestCall(config.localeHost, config.localePath).then((respData) => {
        return respData;
    }).then((resp) => {
        let jsonData = JSON.parse(resp.body);
        //console.log(jsonData);
        return res.json(jsonData);
    }).catch((err) => {
        return res.json(err);
    })
}

exports.listCountries = (req, res) => {
    helper.makeRestCall(config.countriesHost, config.countriesPath).then((respData) => {
        return respData;
    }).then((resp) => {
        let jsonData = JSON.parse(resp.body);
        //console.log(jsonData);
        return res.json(jsonData);
    }).catch((err) => {
        return res.json(err);
    })
}