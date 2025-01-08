function showTab(tab) {
  document.querySelectorAll('.tab-content').forEach((content) => {
    content.style.display = 'none';
  });
  document.getElementById(tab).style.display = 'block';

  document.querySelectorAll('.tabs button').forEach((button) => {
    button.classList.remove('active');
  });
  document.getElementById(`${tab}-tab`).classList.add('active');
}

function convert(type) {
  const value = document.getElementById(`${type}-value`).value;
  const from = document.getElementById(`${type}-from`).value;
  const to = document.getElementById(`${type}-to`).value;

  fetch(`/convert?value=${value}&from=${from}&to=${to}&type=${type}`)
    .then((response) => response.json())
    .then((data) => {
      if (data.error) {
        document.getElementById(`${type}-result`).innerText = `Error: ${data.error}`;
      } else {
        document.getElementById(`${type}-result`).innerText =
          `${value} ${from} = ${data.result} ${to}`;
      }
    })
    .catch((error) => {
      console.error("Error:", error);
      document.getElementById(`${type}-result`).innerText = "Error in conversion!";
    });
}

