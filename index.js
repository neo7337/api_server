const express = require('express');
var bodyParser = require('body-parser');
const app = express();
const cors = require('cors');

app.use(bodyParser.urlencoded({
    extended: true
}));
app.use(bodyParser.json());
app.use(cors());

var port = 8989;

require('./routes/routes')(app);

app.listen(port, function(){
    console.log(`Server is up and listening at ${port}`);
});