let expectedTime = (a, m, b) => Math.floor((a + (4 * m) + b) / 6);
let projectObj = {};
let alphabet = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"]

$(".add-phase").click(function () {
  let addPhase = ($("#phaseName").val()).trim()
let addProj = ($("#projName").val()).trim()
let addAct = ($("#actName").val()).trim()
let addSubAct = ($("#subactName").val()).trim()
  let activityNo = convertToRoman($(".table-body > div").length + 1);
  // ! Object Creator (Done)
  if (addProj == null || addProj == "" || addProj == undefined) {
    alert("Please put a project name.");
    return
  }
  projectObj[addProj] = { ...projectObj[addProj] };
  if (addPhase == null || addPhase == "" || addPhase == undefined) {
    alert("Please put a phase name.");
    return
  }
  if (addAct == null || addAct == "" || addAct == undefined) {
    alert("Please put a activity name.");
    return
  }
  if (addSubAct == null || addSubAct == "" || addSubAct == undefined) {
    alert("Please put a sub-activity name.");
    return
  }
  projectObj[addProj][addPhase] = { ...projectObj[addProj][addPhase] };
  projectObj[addProj][addPhase][addAct] = { ...projectObj[addProj][addPhase][addAct] };
  projectObj[addProj][addPhase][addAct][addSubAct] = addSubAct;
  // projectObj[addProj][addPhase] = { ...projectObj[addProj][addPhase] };
  // if (addAct == null || addAct == "" || addAct == undefined) {
  //   if (addSubAct != null && addSubAct != "" && addSubAct != undefined) {
  //     projectObj[addProj][addPhase][addSubAct] = addSubAct
  //   }
  // }
  // else {
  //   projectObj[addProj][addPhase][addAct] = { ...projectObj[addProj][addPhase][addAct] };
  //   if (addSubAct != null && addSubAct != "" && addSubAct != undefined) {
  //     projectObj[addProj][addPhase][addAct][addSubAct] = addSubAct;
  //   }
  // }
  // Done
  console.log(projectObj)
  projectObj = {
    "Project 1": {
      "Phase 1": {
        "Act 1": {
          "Sub Act 1": "Sub Act 1"
        },
        "Act 2": {
          "Sub Act 1": "Sub Act 1",
          "Sub Act 2": "Sub Act 2",
          "Sub Act 3": "Sub Act 3"
        },
      },
      "Phase 2": {
        "Act 1": {
          "Sub Act 1": "Sub Act 1",
          "Sub Act 2": "Sub Act 2",
          "Sub Act 3": "Sub Act 3"
        },
        "Act 2": {
          "Sub Act 1": "Sub Act 1",
          "Sub Act 2": "Sub Act 2"
        }
      }
    }
  }
  // Table Creator (Under Construction) (DONE?)
  $(".table-body").empty();
  $(".content-from-table").empty();
  $(".totalPricing").remove();
  projectLen = Object.keys(projectObj[addProj]).length
  for (let i = 0; i < projectLen; i++) {
    phaseKey = Object.keys(projectObj[addProj])[i]
    var phaseFormat =
      `<div class="phase-activity-container-${i + 1}">
    <div class="table-content text-center text-break">
      <div>${convertToRoman(i + 1)}.</div>
      <div>${phaseKey}</div>
      <div>--</div>
      <div>--</div>
      <div>--</div>
      <div>--</div>
    </div>
  </div>`
    $(".table-body").append(phaseFormat)
    for (let j = 0; j < Object.keys(projectObj[addProj][phaseKey]).length; j++) {
      actObj = Object.values(projectObj[addProj])[i]

      actKey = Object.keys(actObj)[j]

      var actFormat =
        `<div class="activity-container-${j + 1}">
    <div class="table-content text-center text-break">
      <div>${convertToRoman(i + 1)}. ${alphabet[j]}.</div>
      <div>${actKey}</div>
      <div>--</div>
      <div>--</div>
      <div>--</div>
      <div>--</div>
    </div>
  </div>`
 
      if (actKey == "") {
        var actFormat =
        `<div class="activity-container-${j + 1}">
    <div class="d-none table-content text-center text-break">
      <div>${convertToRoman(i + 1)}. ${alphabet[j]}.</div>
      <div>${actKey}</div>
      <div><input type="number" value="0" min="0" oninput="stringAct(this)"></div>
      <div><input type="number" value="0" min="0" oninput="stringAct(this)"></div>
      <div><input type="number" value="0" min="0" oninput="stringAct(this)"></div>
      <div>--</div>
    </div>
  </div>`
      }
  //     if (typeof Object.values(actObj)[j] === "string") {
  //       console.log('this is a string')
  //       var phaseDirect =
  //         `<div class="direct-sub-activity-container-${j + 1}">
  //   <div class="table-content text-center text-break" style="color:purple;">
  //     <div>${convertToRoman(i + 1)}. ${alphabet[j]}.</div>
  //     <div>${actKey}</div>
  //     <div><input type="number" value="0" min="0" oninput="stringAct(this)"></div>
  //     <div><input type="number" value="0" min="0" oninput="stringAct(this)"></div>
  //     <div><input type="number" value="0" min="0" oninput="stringAct(this)"></div>
  //     <div>--</div>
  //   </div>
  // </div>`
  //       $(`.phase-activity-container-${i + 1}`).append(phaseDirect)
  //       continue
  //     }
      $(`.phase-activity-container-${i + 1}`).append(actFormat)
      for (let k = 0; k < Object.keys(projectObj[addProj][phaseKey][actKey]).length; k++) {
        //actKey = Object.keys(projectObj[addProj][phaseKey])[j]
        subActKey = Object.values(projectObj[addProj][phaseKey][actKey])[k]
        var subActFormat =
          `<div class="sub-activity-container-${k + 1}">
      <div class="table-content text-center text-break">
      <div>${convertToRoman(i + 1)}. ${alphabet[j]}. ${convertToRoman(k + 1).toLowerCase()}.</div>
      <div>${subActKey}</div>
      <div><input type="number" value="0" min="0" oninput="unexpectedTimeRow(this)"></div>
      <div><input type="number" value="0" min="0" oninput="unexpectedTimeRow(this)"></div>
      <div><input type="number" value="0" min="0" oninput="unexpectedTimeRow(this)"></div>
      <div>--</div>
    </div>
    </div>`
    if (subActKey == "") {
      var subActFormat =
          `<div class="sub-activity-container-${k + 1}">
      <div class="d-none table-content text-center text-break">
      <div>${convertToRoman(i + 1)}. ${alphabet[j]}. ${convertToRoman(k + 1).toLowerCase()}.</div>
      <div>${subActKey}</div>
      <div><input type="number" value="0" min="0" oninput="unexpectedTimeRow(this)"></div>
      <div><input type="number" value="0" min="0" oninput="unexpectedTimeRow(this)"></div>
      <div><input type="number" value="0" min="0" oninput="unexpectedTimeRow(this)"></div>
      <div>--</div>
    </div>
    </div>`
    }
        $(`.phase-activity-container-${i + 1} .activity-container-${j + 1}`).append(subActFormat);
      }
    }
  }

  projectLen = Object.keys(projectObj[addProj]).length
  for (let i = 0; i < projectLen; i++) {
    format =
      `<div class="border border-light p-2">
      <div class="container">
      <h3>Activity: ${Object.keys(projectObj[addProj])[i]}</h3>
      <h3>Labor</h3>
        <table class="table" id="laborTable">
            <thead>
                <tr>
                    <th>Worker</th>
                    <th>Quantity</th>
                    <th>Rate</th>
                    <th>Total</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td><input type="text" name="worker[]"></td>
                    <td><input type="number" name="quantity[]" oninput="calculateTotal(this)"></td>
                    <td><input type="number" name="rate[]" oninput="calculateTotal(this)"></td>
                    <td>₱<span class="total">0</span></td>
                </tr>
            </tbody>
        </table>
        <button class="add-row-btn" onclick="addRow(this)">Add Row</button>
    </div>

    <div class="container">
        <h3>Equipment</h3>
        <table class="table" id="equipmentTable">
            <thead>
                <tr>
                    <th>Equipment</th>
                    <th>Quantity</th>
                    <th>Rate</th>
                    <th>Total</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td><input type="text" name="equipment[]"></td>
                    <td><input type="number" name="quantity[]" oninput="calculateTotal(this)"></td>
                    <td><input type="number" name="rate[]" oninput="calculateTotal(this)"></td>
                    <td>₱<span class="total">0</span></td>
                </tr>
            </tbody>
        </table>
        <button class="add-row-btn" onclick="addRow(this)">Add Row</button>
    </div>
  </div>`
    $(".content-from-table").append(format)
  }
  $(".pricing").append(
    `<div class="totalPricing p-3 sticky-top text-center">
      <h5>Live Price</h5>
      <h5>Total Amount: ₱<span id="totalAmount">0</span></h5>
    </div>`)
  // DONE?
  $("#projName").prop("readonly", true)
})
// 

//Total Price Table for each Activities

//function to change values on input change for sub activity which are under phase and are not under activity 
//or sub activity which is in a level of activity
function stringAct(element) {
  var row = element.parentNode.parentNode
  var ot = row.children[2].children[0].value
  var mt = row.children[3].children[0].value
  var pt = row.children[4].children[0].value
  row.children[5].textContent = expectedTime(Number(ot), Number(mt), Number(pt))

  var phase = row.parentNode.parentNode
  var otPhase = 0;
  var mtPhase = 0;
  var ptPhase = 0;
  for (let i = 1; i < phase.children.length; i++) {
    var forOtPhase = phase.children[i].children[0].children[2]
    var forMtPhase = phase.children[i].children[0].children[3]
    var forPtPhase = phase.children[i].children[0].children[4]

    forOtPhase = (forOtPhase.textContent.length != 0) ? forOtPhase.textContent : forOtPhase.children[0].value;
    forMtPhase = (forMtPhase.textContent.length != 0) ? forMtPhase.textContent : forMtPhase.children[0].value;
    forPtPhase = (forPtPhase.textContent.length != 0) ? forPtPhase.textContent : forPtPhase.children[0].value;

    otPhase += Number(forOtPhase)
    mtPhase += Number(forMtPhase)
    ptPhase += Number(forPtPhase)
  }
  phase.children[0].children[2].textContent = !isNaN(otPhase) ? otPhase : "--";
  phase.children[0].children[3].textContent = !isNaN(mtPhase) ? mtPhase : "--";
  phase.children[0].children[4].textContent = !isNaN(ptPhase) ? ptPhase : "--";

  phase.children[0].children[5].textContent = !isNaN(expectedTime(otPhase, mtPhase, ptPhase)) ? expectedTime(otPhase, mtPhase, ptPhase) : "--";
}

//function to change values on input change for sub activity
function unexpectedTimeRow(element) {
  // Child Sub Act Container
  var subActChild = element.parentNode.parentNode;
  var ot = subActChild.children[2].children[0].value;
  var mt = subActChild.children[3].children[0].value;
  var pt = subActChild.children[4].children[0].value;
  subActChild.children[5].textContent = expectedTime(Number(ot), Number(mt), Number(pt));

  //Child of Act Container
  var actChild = subActChild.parentNode.parentNode;
  var otAct = 0;
  var mtAct = 0;
  var ptAct = 0;
  var etAct = 0;
  for (let i = 1; i < actChild.children.length; i++) {
    otAct += Number(actChild.children[i].children[0].children[2].children[0].value);
    mtAct += Number(actChild.children[i].children[0].children[3].children[0].value);
    ptAct += Number(actChild.children[i].children[0].children[4].children[0].value);
    etAct += Number(actChild.children[i].children[0].children[5].textContent);
  }

  actChild.children[0].children[2].textContent = otAct;
  actChild.children[0].children[3].textContent = mtAct;
  actChild.children[0].children[4].textContent = ptAct;
  actChild.children[0].children[5].textContent = !isNaN(etAct) ? etAct : "--";

  //Phase Container
  var phase = actChild.parentNode
  var otPhase = 0;
  var mtPhase = 0;
  var ptPhase = 0;
  for (let i = 1; i < phase.children.length; i++) {
    var forOtPhase = phase.children[i].children[0].children[2]
    var forMtPhase = phase.children[i].children[0].children[3]
    var forPtPhase = phase.children[i].children[0].children[4]

    forOtPhase = (forOtPhase.textContent.length != 0) ? forOtPhase.textContent : forOtPhase.children[0].value;
    forMtPhase = (forMtPhase.textContent.length != 0) ? forMtPhase.textContent : forMtPhase.children[0].value;
    forPtPhase = (forPtPhase.textContent.length != 0) ? forPtPhase.textContent : forPtPhase.children[0].value;

    otPhase += Number(forOtPhase)
    mtPhase += Number(forMtPhase)
    ptPhase += Number(forPtPhase)
  }
  phase.children[0].children[2].textContent = !isNaN(otPhase) ? otPhase : "--";
  phase.children[0].children[3].textContent = !isNaN(mtPhase) ? mtPhase : "--";
  phase.children[0].children[4].textContent = !isNaN(ptPhase) ? ptPhase : "--";

  phase.children[0].children[5].textContent = !isNaN(expectedTime(otPhase, mtPhase, ptPhase)) ? expectedTime(otPhase, mtPhase, ptPhase) : "--";
}

function calculateTotal(input) {
  const row = input.closest('tr');
  const quantity = row.querySelector('input[name="quantity[]"]').value;
  const rate = row.querySelector('input[name="rate[]"]').value;
  const total = quantity * rate;
  row.querySelector('.total').textContent = total;
  calculateGrandTotal();
}

function calculateGrandTotal() {
  let grandTotal = 0;
  document.querySelectorAll('.total').forEach(function (element) {
    grandTotal += parseFloat(element.textContent);
  });
  document.getElementById('totalAmount').textContent = grandTotal;
}

function addRow(element) {
  const table = element.parentNode;
  const tbody = table.querySelector('tbody');
  const newRow = document.createElement('tr');
  newRow.innerHTML = `
        <td><input type="text" name="worker[]"></td>
        <td><input type="number" name="quantity[]" oninput="calculateTotal(this)"></td>
        <td><input type="number" name="rate[]" oninput="calculateTotal(this)"></td>
        <td>₱<span class="total">0</span></td>
    `;
  tbody.appendChild(newRow);
}

function convertToRoman(num) {
  //values of the numbers
  let numerals = {
    M: 1000,
    CM: 900,
    D: 500,
    CD: 400,
    C: 100,
    XC: 90,
    L: 50,
    XL: 40,
    X: 10,
    IX: 9,
    V: 5,
    IV: 4,
    I: 1,
  };
  //passing the new values of the new numbers to convert into numerals
  let newNumeral = "";
  let numCopy = num;   //copy the num so it doesn't mutate and can used inside the for loop

  //checking the values of numerals objects
  for (let i in numerals) {
    //using the j variable of num, if the num is still greater than the numerals index 
    //it will increment the key to the variable newNumerals and stop the loop,
    // it will subtract the num to the values of numerals .
    for (let j = 0; j <= num; j++) {
      if (numCopy >= numerals[i]) {
        newNumeral += i;
        numCopy -= numerals[i];
      }
    }
  }
  //return to newNumerals to see the new value.
  return newNumeral;
}

// TODO: Schedule Completion
// Save plan which is next page

$(".save-plan").click(function () {
  var completeSchedObj = {}
  let phases = $(".table-body").children()
  for (var i = 0; i < phases.length; i++) {
    var phaseNum = $(phases[i].children[0].children[0]).text()
    var phaseVal = $(phases[i].children[0].children[1]).text()
    var phaseOt = !isNaN(Number($(phases[i].children[0].children[2]).text())) ? Number($(phases[i].children[0].children[2]).text()) : 0;
    var phaseMt = !isNaN(Number($(phases[i].children[0].children[3]).text())) ? Number($(phases[i].children[0].children[3]).text()) : 0;
    var phasePt = !isNaN(Number($(phases[i].children[0].children[4]).text())) ? Number($(phases[i].children[0].children[4]).text()) : 0;
    var phaseobj = { [phaseNum]: { [phaseVal]: [phaseOt, phaseMt, phasePt] } }
    Object.assign(completeSchedObj, phaseobj)
    var activities = $(phases[i]).children()
    for (let j = 1; j < activities.length; j++) {
      var actNum = $(activities[j].children[0].children[0]).text()
      var actVal = $(activities[j].children[0].children[1]).text()
      var actOt = !isNaN(Number($(activities[j].children[0].children[2]).text())) ? Number($(activities[j].children[0].children[2]).text()) : 0;
      var actMt = !isNaN(Number($(activities[j].children[0].children[3]).text())) ? Number($(activities[j].children[0].children[3]).text()) : 0;
      var actPt = !isNaN(Number($(activities[j].children[0].children[4]).text())) ? Number($(activities[j].children[0].children[4]).text()) : 0;
      var actobj = { [actNum]: { [actVal]: [actOt, actMt, actPt] } }
      Object.assign(phaseobj[phaseNum], actobj)
      var subActivities = $(activities[j]).children()
      for (let k = 1; k < subActivities.length; k++) {
          var subActivitiesKey = $(subActivities[k].children[0].children[0]).text();
          var subActivitiesVal = $(subActivities[k].children[0].children[1]).text();
          var subActOt = !isNaN(Number($(subActivities[k].children[0].children[2].children[0]).val())) ? Number($(subActivities[k].children[0].children[2].children[0]).val()) : 0;
          var subActMt = !isNaN(Number($(subActivities[k].children[0].children[3].children[0]).val())) ? Number($(subActivities[k].children[0].children[3].children[0]).val()) : 0;
          var subActPt = !isNaN(Number($(subActivities[k].children[0].children[4].children[0]).val())) ? Number($(subActivities[k].children[0].children[4].children[0]).val()) : 0;
          var subactobj = { [subActivitiesKey]: { [subActivitiesVal]: [subActOt, subActMt, subActPt] } }
          Object.assign(phaseobj[phaseNum][actNum], subactobj)
        } 
    }
  }
  // console.log(completeSchedObj)
  $.ajax({
    url:"/api/add-project",
    method:"POST",
    contentType: "application/json",
    data: JSON.stringify({
      projectName: ($("#projName").val()).trim(),
      cost: Number($("#totalAmount").text()),
      obj: completeSchedObj,
    }),
    success: function(resp) { 
      console.log(resp["response"])
      location.href = location.origin + "/user/schedule-completion"
    },
    error: function(response) {
      console.log(response["responseText"])
    }
  })
})
