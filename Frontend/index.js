let url = "http://localhost:8080/books";

let tbody = document.querySelector("tbody"),
  form = document.querySelector(".form"),
  inputs = document.querySelectorAll(".input_box_form input");

async function fetchData() {
  try {
    const response = await fetch(url, { method: "GET" });
    const data = await response.json();
    innerToTable(data);
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

fetchData()
  .then((data) => {
    console.log(data);
  })
  .catch((error) => {
    console.error(error);
  });

function innerToTable(data) {
  tbody.innerHTML = "";

  if (data.length) {
    data.map((el, index) => {
      tbody.innerHTML += `
            <tr>
                <td>${index + 1}</td>
                <td class="title_td">${el?.title}</td>
                <td class="author_td">${el?.author}</td>
                <td>${el?.quantity}</td>
                <td class="edit_td" onclick="handleEdit()">Edit</td>
                <td class="delete_td">
                  <button onclick="handleDelete('${el.Id}')">Delete</button>
                </td>
            </tr>
        `;
    });
  } else {
    tbody.innerHTML = `
      <tr>
        <td colspan="100" class="no_data" rowspan="1">
            <b>No data</b>
            <span>
                <p>.</p>
                <p>.</p>
                <p>.</p>
            </span>
      </td>
    </tr>
    `;
  }
}

let inputsValue = {
  title: "",
  author: "",
  quantity: 0,
};

function clearInputValue() {
  inputs.forEach((el) => {
    el.value = "";
  });
  inputsValue = {
    title: "",
    author: "",
    quantity: 0,
  };
}

form.addEventListener("submit", (event) => {
  event.preventDefault();

  inputs.forEach((el) => {
    inputsValue = {
      ...inputsValue,
      [el.name]: el.value,
    };
  });

  postData(inputsValue);
});

async function postData(obj) {
  console.log(obj);
  console.log(JSON.stringify(obj));
  try {
    fetch(url, {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...obj, quantity: +obj.quantity }),
    }).then((res) => fetchData(url));
  } catch (error) {
    console.error(error);
  }
  clearInputValue();
}

async function handleDelete(id) {
  try {
    let response = await fetch(url + `/${id}`, { method: "DELETE" });
    console.log(response);
    if (response.ok) {
      console.log("Item deleted successfully");
      fetchData();
    } else {
      console.log("Error deleting item:", response.statusText);
    }
  } catch (error) {
    console.error("An error occurred during delete operation:", error);
  }
}

async function handleEdit() {
  console.log("Salom");
}
