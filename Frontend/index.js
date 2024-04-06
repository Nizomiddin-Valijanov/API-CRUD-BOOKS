let url = "http://localhost:8080/books";

let tbody = document.querySelector("tbody"),
  form = document.querySelector(".form"),
  inputs = document.querySelectorAll(".input_box_form input"),
  submitBtn = document.querySelector(".submit_btn");

let update = false;

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
                <td class="index_td">${index + 1}</td>
                <td class="title_td">${el?.title}</td>
                <td class="author_td">${el?.author}</td>
                <td>${el?.quantity}</td>
                <td class="edit_td">
                  <button onclick="handleEdit('${el.Id}')">Edit</button>
                </td>
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

  if (update) {
    patchData(inputsValue);
  } else {
    postData(inputsValue);
  }
});

async function patchData(obj) {
  try {
    const response = await fetch(url + `/${obj.Id}`, {
      method: "PUT",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(obj),
    });

    if (!response.ok) {
      throw new Error("Failed to patch data");
    }

    clearInputValue();
  } catch (error) {
    console.error(error);
  }
}

async function postData(obj) {
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

async function handleEdit(id) {
  try {
    const response = await fetch(url + `/${id}`, { method: "GET" });
    inputsValue = await response.json();
  } catch (error) {
    console.error(error);
  }

  inputs.forEach((input) => {
    input.value = inputsValue[input.name];
  });
  submitBtn.innerHTML = "Update";
  update = true;
}
