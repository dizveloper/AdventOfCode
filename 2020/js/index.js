const one = require('./days/day-one');
// const two = require('./days/day-two');
// const three = require('./days/day-three');
// const four = require('./days/day-four');
// const five = require('./days/day-five');
// const six = require('./days/day-six');
// const seven = require('./days/day-seven');
// const eight = require('./days/day-eight');
// const nine = require('./days/day-nine');
// const ten = require('./days/day-ten');
// const eleven = require('./days/day-eleven');
// const twelve = require('./days/day-twelve');
// const thirteen = require('./days/day-thirteen');
// const fourteen = require('./days/day-fourteen');
// const fifteen = require('./days/day-fifteen');
// const sixteen = require('./days/day-sixteen');
// const seventeen = require('./days/day-seventeen');
// const eighteen = require('./days/day-eighteen');
// const nineteen = require('./days/day-nineteen');
// const twenty = require('./days/day-twenty');
// const twentyOne = require('./days/day-twenty-one');
// const twentyTwo = require('./days/day-twenty-two');
// const twentyThree = require('./days/day-twenty-three');
// const twentyFour = require('./days/day-twenty-four');
// const twentyFive = require('./days/day-twenty-five');





const args = process.argv.slice(2)

const whichDay = (which) => {
	switch (which) {
	case "1":
		console.log("Running day one.")
        one.one()
        break;
	case "2":
        console.log("Running day two.")
        two.two()
        break;
	case "3":
        console.log("Running day three.")
        three.three()
        break;
	case "4":
        console.log("Running day four.")
        four.four()
        break;
	case "5":
        console.log("Running day five.")
        five.five()
        break;
	case "6":
        console.log("Running day six.")
        six.six()
        break;
	case "7":
        console.log("Running day seven.")
        seven.seven()
        break;
	case "8":
        console.log("Running day eight.")
        eight.eight()
        break;
	case "9":
        console.log("Running day nine.")
        nine.nine()
        break;
	case "10":
        console.log("Running day ten.")
        ten.ten()
        break;
	case "11":
        console.log("Running day eleven.")
        eleven.eleven()
        break;
	case "12":
        console.log("Running day twelve.")
        twelve.twelve()
        break;
	case "13":
        console.log("Running day thirteen.")
        thirteen.thirteen()
        break;
	case "14":
        console.log("Running day fourteen.")
        fourteen.fourteen()
        break;
	case "15":
        console.log("Running day fifteen.")
        fifteen.fifteen()
        break;
	case "16":
        console.log("Running day sixteen.")
        sixteen.sixteen()
        break;
	case "17":
        console.log("Running day seventeen.")
        seventeen.seventeen()
        break;
	case "18":
        console.log("Running day eighteen.")
        eighteen.eighteen()
        break;
	case "19":
        console.log("Running day nineteen.")
        nineteen.nineteen()
        break;
	case "20":
        console.log("Running day twenty.")
        twenty.twenty()
        break;
	case "21":
        console.log("Running day twenty one.")
        twentyOne.twentyOne()
        break;
	case "22":
        console.log("Running day twenty two.")
        twentyTwo.twentyTwo()
        break;
	case "23":
        console.log("Running day twenty three.")
        twentyThree.twentyThree()
        break;
	case "24":
        console.log("Running day twenty four.")
        twentyFour.twentyFour()
        break;
	case "25":
        console.log("Running day twenty five.")
        twentyFive.twentyFive()
        break;
	default:
		console.log("Pick a day between 1 and 25")
	}
}

whichDay(args[0])