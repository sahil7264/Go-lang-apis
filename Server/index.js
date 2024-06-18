const express = require('express')

const cors = require('cors')

const app = express()

app.use(express.json())
app.use(cors())
app.get('/', (req, res) => {
    res.send({
        message: "Hey Sahil, Hello from Server"
    })
});
app.post('/post', (req, res) => {
    let { a, b } = req.body

    a = parseInt(a)
    b = parseInt(b)
    res.send({
        sum : a + b  
    })
});
app.post('/form', (req, res) => {
    console.log(req.data)
    let { fname, lname,mname } = req.body
    console.log(fname,lname,mname);
    res.send({
        message : "Ok"
    })
});

app.listen(4000, () => {
    console.log("Server listening on port 4000")
}) 