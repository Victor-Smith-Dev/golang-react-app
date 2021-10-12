import React from "react";
import { useProducts } from "../hooks/useProducts";
import { Product } from "../interfaces/Response";
import { Link } from "react-router-dom";

export const HomeScreen = () => {
  const { productos } = useProducts(); 

  const renderItem = ({ name, ID, price }: Product) => {
    return (
      <tr key={ID.toString()}>
        <td>{ID}</td>
        <td>{name}</td>
        <td>{price}</td>
        <td>              
          <Link to={`/historyPrices/${ID}`}>Historial</Link>
        </td>
      </tr>
    );
  };

  return (
    <>
      <h1>Priceomatic Client</h1>
      <hr />
      <div style={{
        height: '300px',
        overflowY: 'scroll',
      }}>
        <h3>Productos: ({productos.length})</h3>
        <table className="table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Nombre</th>
              <th>Precio</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>{productos.map((product: Product) => renderItem(product))}</tbody>
        </table>
      </div>
    </>
  );
};
