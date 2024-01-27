import React from 'react';

function ButtonComponent() {
  // Function to handle button click and hit the endpoint
  const handleViewClick = async () => {
    try {
      const response = await fetch('http://localhost:8084/view', {
  method: 'GET',
});

if (!response.ok) {
  throw new Error('Network response was not ok');
}

const data = await response.text(); // Get the plain text response
console.log(data); 

    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  return (
    <div>
      <button onClick={handleViewClick}>View</button>
      <button>Upload</button>
      <button>Download</button>
      <button>Delete</button>
    </div>
  );
}

export default ButtonComponent;
