// Helper for displaying status messages.
export const addMessage = (message) => {
  const messagesDiv = document.querySelector('#card-messages');
  messagesDiv.classList.remove("hidden")
  const messageWithLinks = addDashboardLinks(message);
  messagesDiv.innerHTML += `${messageWithLinks}<br>`;
  console.log(`Debug: ${message}`);
};

// Adds links for known Stripe objects to the Stripe dashboard.
export const addDashboardLinks = (message) => {
  const piDashboardBase = 'https://dashboard.stripe.com/test/payments';
  return message.replace(
    /(pi_(\S*)\b)/g,
    `<a href="${piDashboardBase}/$1" target="_blank">$1</a>`
  );
};
