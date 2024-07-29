
//<script src="https://js.stripe.com/v3/"></script>
// import {loadStripe} from '../node_modules/@stripe/stripe-js'

import { loadConnectAndInitialize } from "https://cdn.jsdelivr.net/npm/@stripe/connect-js/+esm";
import * as stripe from "https://cdn.jsdelivr.net/npm/@stripe/stripe-js/+esm";
 const stripeins = await stripe.loadStripe('pk_test_51PDPzU1UayZGZkAxOTvpNoa6qOgLzSj7lzyHNtyq1gJvfjSqSrIF0Lselt18fRgApbeHuE2WmqARc3KVKdSl7o4d00V0PI5lE3');

async function initialize() {


      const connectedAccountId = await fetchConnectedAccountId();
//
    if (!connectedAccountId) {
        return;
    }
//
  const fetchClientSecret = async () => {
    const response = await fetch("/checkout/session", {
      method: "POST",
      body:
          JSON.stringify({
    account: connectedAccountId,

    }),
    });
    const { client_secret } = await response.json();
    console.log(client_secret)
    return client_secret;
  };

  // Initialize Checkout
  const checkout = await stripeins.initEmbeddedCheckout({
    fetchClientSecret,
  });

  // Mount Checkout
  checkout.mount('#checkout');
}
async function fetchConnectedAccountId() {
    try {
        const response = await fetch("/billing/account", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            }
        });
        const { account, error } = await response.json();

        if (error) {
            showError();
            return null;
        }

        return account;
    } catch (err) {
        console.error("Error fetching connected account ID:", err);
        showError();
        return null;
    }
}

function showError() {
    document.getElementById("error").classList.remove("hidden");
    document.getElementById("sign-up-button").classList.remove("hidden");
    document.getElementById("creating-connected-account").classList.add("hidden");
    document.getElementById("dev-callout").classList.add("hidden");
}

initialize();
