//document.domain = 'localhost';

var app = new Vue ({
    el: '#app',
    data: {
        message: 'Hello Vue!',
        groceryList: [
            { text: "Vegetables" },
            { text: "Cheese" }
        ],
        yesno: 'not set',
        styleObject: {
            color: 'red',
            fontSize: '10em',
            
        }
    },
    mounted () {
        console.log("in ready")

        this.$http.get("http://localhost:8080").then(response => {
            this.yesno = response.body;
        }, response => {
            console.log("Error " + response)
        });
    }
});


//Vue.http.options.xhr = {withCredentials: true}
//Vue.http.options.root = 'localhost:8080'
console.log("done")