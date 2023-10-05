import Card from 'react-bootstrap/Card';
import './CardComponent.css';
import CardComponent from './CardComponent'

function WebPage() {
  return (
   <div className="container">
    <header>
        <h1 align="center">Chaos Mesh Report</h1>
        <hr />

                
        <div className="logsContainer">
        <Card className='card0'>
            <Card.Body>
                <p>TiDB for PG Latest</p>
            </Card.Body>
        </Card>
            <CardComponent />
            {/* <CardComponent /> */}
            {/* <CardComponent /> */}
        </div>
        <hr />


    </header>
   </div>
  );
}

export default WebPage;
