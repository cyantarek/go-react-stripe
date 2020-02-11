import React, {useState} from 'react';
import './App.css';
import StripeCheckout from "react-stripe-checkout";

function App() {
    const [total, setTotal] = useState(0);

    const handleChange = (e) => {
        setTotal(e.target.value)
    };

    const handleStripeToken = (token) => {
        fetch("http://localhost:5300/api/payment", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                token,
                amount: total * 100
            })
        }).then(res => {
            if (res.status !== 200) {
                throw res.status
            } else {
                return res.json
            }
        })
            .then(data => {
                console.log(data)
            })
    };

    return (
        <div className="App">
            <div className="card">
                <div className="content">
                    <input onChange={handleChange} className={"input_field"} type="text"
                           placeholder={"amount"}/>
                    <StripeCheckout className={"stripe__button"}
                                    amount={total * 100}
                                    token={handleStripeToken}
                                    label={"Pay with Stripe"}
                                    stripeKey="pk_test_xdfOMo2yF5PxOKXF8DTY39QE"
                    />
                </div>
            </div>
        </div>
    );
}

export default App;
