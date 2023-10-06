import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';

function CardComponent() {
  // Create a new Date object
  const today = new Date();

  // Get the year, month, and day
  const year = today.getFullYear();
  const month = today.getMonth() + 1; // Months are zero-based, so add 1
  const day = today.getDate();

  // Format the date as a string (e.g., "YYYY-MM-DD")
  const formattedDate = `/pdfs/report-${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}.pdf`;

  console.log(formattedDate);
  return (
    <Card className='cardStyle'>
      <Card.Body>
        <Button variant="light">Date</Button>
        <Button variant="light">
          <a href={formattedDate} target="_blank" rel="noreferrer">Check Report</a>
        </Button>
      </Card.Body>
    </Card>
  );
}

export default CardComponent;