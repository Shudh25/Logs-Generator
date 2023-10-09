// Create a new Date object
const today = new Date();

// Get the year, month, and day
const year = today.getFullYear();
const month = today.getMonth() + 1; // Months are zero-based, so add 1
const day = today.getDate();
const hour = today.getHours();

// Format the date as a string (e.g., "YYYY-MM-DD")
// const formattedDate = `pdfs/report-${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}.pdf`;
const formattedDate = `pdfs/report-${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}-${hour.toString().padStart(2, '0')}.pdf`;

console.log(formattedDate);

// Create a new anchor element
const link = document.createElement('a');

// Set the URL (href) using setAttribute
link.setAttribute('href', formattedDate);

// Set the link text
link.textContent = 'Check Report';

// Append the anchor element to the document or another HTML element
document.getElementById('btn').appendChild(link); // Appending it to the body for this example
