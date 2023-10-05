import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';

function CardComponent() {
  return (
    <Card className='cardStyle'>
      <Card.Body>
        <Button variant="light">Date</Button>
        <Button variant="light">
          <a href="/pdfs/report-2023-10-06.pdf" target="_blank">Check Report</a>
        </Button>
      </Card.Body>
    </Card>
  );
}

export default CardComponent;