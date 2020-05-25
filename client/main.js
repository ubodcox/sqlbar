folder = 'tasks';
var task_id = 1

setResultTableVisible(false);
sendData('/api/v1/tasks/get', refreshTaskList);

//--- Check -------------------------------------------------------------------

function sendTaskCheck() {
    div = document.getElementById('sql');
    if (!div) {
        console.log("no sql element")
        return
    }

    sql = div.value
    console.log('sql: ' + sql)
    sendData('/api/v1/tasks/check?id=' + task_id + '&sql=' + sql, refreshTaskCheck);
}

function refreshTaskCheck(response) {
    console.log("refreshCheck", response)
    element = response

    div = document.getElementById('answer');
    div.innerHTML = element["answer"];    

    refreshResultTable(response)
}

function  refreshTaskList(response) {
    grid = document.getElementById('ttasks')
    tbody = grid.getElementsByTagName('tbody')
    console.log("tbody", tbody)

    var list = response.tasks;
    console.log('list: ', list)
    //const thead = document.createElement("thead");
    //element = list[0];
    //console.log('element: ', element)
    //addHeaderLine(thead, element);
    //grid.appendChild(thead)

    //const tbody = document.createElement("tbody")
    //grid.appendChild(tbody)
    for (var i = 0; i < list.length; i++) {
        element = list[i];
        const trow = document.createElement("tr");
        addNewTaskLine(trow, element);
        console.log("tbody", tbody)
        trow.className = "grid_task_tr"

        trow.onclick = function() { selectTask(element["id"]) } ;

        grid.appendChild(trow)
    };
    if (list.length > 0) {
        element = list[0];
        selectTask(element["id"]);
    }
}

function addNewTaskLine(trow, element) {
    th = document.createElement("th");
    th.innerHTML = element["stars"];
    trow.appendChild(th);
    
    th = document.createElement("th");
    th.innerHTML = element["text"].substring(0, 20)+"...";
    trow.appendChild(th);
}

//-- Get ----------------------------------------------------------------------

function refreshTaskGet(response) {
    console.log("refreshGet", response)
    element = response["tasks"][0]

    div = document.getElementById('id');
    div.innerHTML = element["id"];

    div = document.getElementById('stars');
    div.innerHTML = element["stars"];

    div = document.getElementById('task');
    div.innerHTML = element["text"];

    div = document.getElementById('cols');
    div.innerHTML = element["cols"];

    div = document.getElementById('answer');
    div.innerHTML = '' //element["answer"];
}

function refreshResultTable(result) {
    console.log("refreshResultTable: ", result)
    tgrid = document.getElementById('tgrid')
    /*console.log("grid1: ", grid.innerHTML)
    grid.innerHTML = ""
    console.log("grid2: ", grid.innerHTML)*/

    while (tgrid.firstChild) {
        tgrid.removeChild(tgrid.firstChild);
    }

    if (!result) {
        setResultTableVisible(false)
        return
    }

    if (result["code"] < 0) {
        setResultTableVisible(false)
        return
    }    

    setResultTableVisible(true)

    var list = result.result.list;
    console.log('list: ', list)
    const thead = document.createElement("thead");
    element = list[0];
    console.log('element: ', element)
    addResultHeaderLine(thead, element);
    tgrid.appendChild(thead)

    const tbody = document.createElement("tbody")
    tgrid.appendChild(tbody)
    for (var i = 0; i < list.length; i++) {
        element = list[i];
        const trow = document.createElement("tr");
        addResultNewLine(trow, element);
        trow.className = "grid_task_tr"
        tbody.appendChild(trow)
    };    
}

function addResultHeaderLine(thead, element) {
    var tr = document.createElement("tr"); 
    for (var key of Object.keys(element)) {
        th = document.createElement("th");
        th.innerHTML = key;
        tr.appendChild(th);
    }
    thead.appendChild(tr);
}

function addResultNewLine(trow, element) {
    for (var value of Object.values(element)) {
        th = document.createElement("th");
        th.innerHTML = value;
        trow.appendChild(th);
    }
};


function setResultTableVisible(visible) {
    grid = document.getElementById('tgrid_back')
    if (visible) {
        grid.style.display = "inline";
    } else {
        grid.style.display = "none";
    }
}

//-- Select -------------------------------------------------------------------

function selectTask(id) {
    console.log("selectTask " + id)
    setResultTableVisible(false)

    task_id = id;
    sendData('/api/v1/tasks/get?id=' + task_id, refreshTaskGet);
}

//-- send ---------------------------------------------------------------------

function sendData(url, refreshTable) {
    console.log("refreshData", url)
    try {
        var xhttp = new XMLHttpRequest();
        xhttp.open("GET", url, false);
        xhttp.setRequestHeader("Content-type", "text/html");
        xhttp.send();
    } catch (error) {
        console.log("ERROR1: " + error.message);
    }

    if (xhttp.response[0] == '{') {
        //console.log(xhttp.response)
        response = JSON.parse(xhttp.response);
        console.log("response ", response)
        if (response["code"] < 0) {
            console.log("toastr message: ", response["msg"])
            //toastr.error(response["msg"]);
        }

        refreshTable(response) 
    }
}
