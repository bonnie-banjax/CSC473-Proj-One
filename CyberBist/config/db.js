const mongoose = require('mongoose');

const connectDB = async () => {
  try {
    const conn = await mongoose.connect(process.env.MONGO_URI);
    console.log(`[SYS_LOG]: MongoDB Connected: ${conn.connection.host}`);
  } catch (error) {
    console.error(`[SYS_ERROR]: ${error.message}`);
    process.exit(1);
  }
};

module.exports = connectDB;