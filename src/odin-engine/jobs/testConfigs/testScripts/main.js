const fs = require("fs");

let writeStream = fs.createWriteStream("output.txt");
writeStream.write("this is a test", "utf-8");
writeStream.on("finish", () => {
    console.log("wrote all data to file");
});
