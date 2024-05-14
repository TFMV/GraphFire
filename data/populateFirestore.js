const admin = require('firebase-admin');

// Load the service account key
const serviceAccount = require('./sa.json');

// Initialize the Firebase Admin SDK
admin.initializeApp({
  credential: admin.credential.cert(serviceAccount),
  databaseURL: `https://${serviceAccount.project_id}.firebaseio.com`,
});

const db = admin.firestore();

// Sample data to populate
const sampleData = {
  users: [
    { id: 'user1', name: 'John Doe', email: 'john@example.com', age: 30 },
    { id: 'user2', name: 'Jane Smith', email: 'jane@example.com', age: 25 },
  ],
  products: [
    { id: 'product1', name: 'Laptop', price: 999.99, stock: 50 },
    { id: 'product2', name: 'Phone', price: 599.99, stock: 100 },
  ],
  orders: [
    { id: 'order1', userId: 'user1', productId: 'product1', quantity: 1, total: 999.99 },
    { id: 'order2', userId: 'user2', productId: 'product2', quantity: 2, total: 1199.98 },
  ],
};

// Function to populate Firestore with sample data
const populateFirestore = async () => {
  try {
    for (const [collection, documents] of Object.entries(sampleData)) {
      const collectionRef = db.collection(collection);
      for (const doc of documents) {
        console.log(`Adding document to ${collection}:`, doc);
        await collectionRef.doc(doc.id).set(doc);
      }
    }
    console.log('Sample data populated successfully.');
  } catch (error) {
    console.error('Error populating Firestore:', error);
  }
};

populateFirestore();
