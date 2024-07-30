import { loadConnectAndInitialize } from "https://cdn.jsdelivr.net/npm/@stripe/connect-js/+esm";

let connectedAccountId = null;

const connectedAccountIdElement = document.getElementById("stripe-account-id");

connectedAccountId = connectedAccountIdElement.innerHTML;

const fetchClientSecret = async () => {
  // Fetch the AccountSession client secret
  const response = await fetch('/billing/session', {
   method: "POST",
    headers: {
    "Content-Type": "application/json",
    },
    body: JSON.stringify({
    account: connectedAccountId,
    }),
  });
  if (!response.ok) {
    // Handle errors on the client side here
    const {error} = await response.json();
    console.log('An error occurred: ', error);
    document.querySelector('#container').setAttribute('hidden', '');
    document.querySelector('#error').removeAttribute('hidden');
    return undefined;
  } else {
    const {client_secret: clientSecret} = await response.json();
    document.querySelector('#container').removeAttribute('hidden');
    document.querySelector('#error').setAttribute('hidden', '');
    return clientSecret;
  }
}

  const instance = loadConnectAndInitialize({
    // This is your test publishable API key.
    publishableKey: "pk_test_51PDPzU1UayZGZkAxOTvpNoa6qOgLzSj7lzyHNtyq1gJvfjSqSrIF0Lselt18fRgApbeHuE2WmqARc3KVKdSl7o4d00V0PI5lE3",
    fetchClientSecret: fetchClientSecret,
    appearance: {
      overlays: 'dialog',
      variables: {
        colorPrimary: '#625afa',
      },
    },
      });
      const container = document.getElementById("container");
      const paymentsComponent = instance.create("payments");
      container.appendChild(paymentsComponent);

      const payoutsContainer = document.getElementById("payouts-container");
      const payouts = instance.create('payouts');
      payoutsContainer.appendChild(payouts);
