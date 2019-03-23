const getFontSize = (textLength) => {
    let fontSize = 27;
    if (textLength >= 5) {
       fontSize -= 10
    }
    return `${fontSize}vw`
};

let grade = document.querySelector("#grade");
grade.style.fontSize = getFontSize(grade.textContent.length);


function zero(a) {
    if (a <= 9) {
        return "0" + a.toString()
    } else {
        return a.toString()
    }
}

let span = document.getElementById('time');

let day = new Date().getDay();
let month = new Date().getMonth();

day = zero(day);
month = zero(month);

let mmdd = month + "/" + day + " ";

function time() {
    let d = new Date();
    let s = zero(d.getSeconds());
    let m = zero(d.getMinutes());
    let h = zero(d.getHours());
    span.textContent = mmdd + h + ":" + m + ":" + s;

    if (m === "11" && s === "00") {
        location.reload();
    }
}

setInterval(time, 1000);