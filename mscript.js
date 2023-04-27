const fs = require("fs");
db = connect("mongodb://192.168.100.239/myDatabase");
console.log(process.argv);
const updateDataToDatabase = async function (filePath, encoding, stockSign) {
  let stockContentReading = fs.readFileSync(filePath, encoding, stockSign);
  if (stockSign === "pre") {
    await db.premarketstocks.insertMany(JSON.parse(preStock));
  }

  if (stockSign === "nif") {
    await db.niftyStocks.insertMany(JSON.parse(stockContentReading));
  }
};
