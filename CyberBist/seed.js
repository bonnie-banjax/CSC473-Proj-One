const mongoose = require('mongoose');
require('dotenv').config();
const MenuItem = require('./models/MenuItem');

const initialMenuItems = [
  {
    name: 'Synth-Wagyu Cuboid',
    price: 18.00,
    description: 'Reconstituted beef proteins suspended in artificial lipid-emulsion on high-density carbohydrate bun. Enhanced with synthetic truffle flavor-code.',
    imageUrl: 'https://images.unsplash.com/photo-1568901346375-23c9450c58cd?auto=format&fit=crop&w=500&q=80',
    category: 'Sludge Allocations'
  },
  {
    name: 'Thermal Dough Disc',
    price: 16.50,
    description: 'Flatbread mechanism loaded with coagulated cow lipids and high-sodium plant extract fluids. Broiled at 450°C inside corporate kilns.',
    imageUrl: 'https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=500&q=80',
    category: 'Sludge Allocations'
  },
  {
    name: 'Heavy-Metal Salmon Base',
    price: 21.00,
    description: 'Slices of aquaculture flesh loaded with Omega-3 and trace ocean microplastics, stacked perfectly upon bleached carbohydrate grains.',
    imageUrl: 'https://images.unsplash.com/photo-1546069901-ba9599a7e63c?auto=format&fit=crop&w=500&q=80',
    category: 'Sludge Allocations'
  }
];

const seedDB = async () => {
  try {
    // 1. Establish database connection
    await mongoose.connect(process.env.MONGO_URI);
    console.log('[SYS_LOG]: Connected to MongoDB for seeding...');

    // 2. Clear existing menu items to prevent duplicates
    await MenuItem.deleteMany({});
    console.log('[SYS_LOG]: Existing menu buffer purged.');

    // 3. Insert initial menu items
    const createdItems = await MenuItem.insertMany(initialMenuItems);
    console.log(`[SYS_LOG]: Successfully ingested ${createdItems.length} menu items into database.`);

    // 4. Close connection cleanly
    await mongoose.connection.close();
    console.log('[SYS_LOG]: Connection terminated. Seeding complete.');
    process.exit(0);
  } catch (error) {
    console.error(`[SYS_ERROR]: Failed to seed database: ${error.message}`);
    process.exit(1);
  }
};

seedDB();
