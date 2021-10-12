import React from "react";
import { useProducts } from "../hooks/useProducts";
import { Product } from "../interfaces/Response";

export const HomeScreen = () => {
  const { productos } = useProducts();

  const renderItem = ({ name, ID, price }: Product) => {
    return (
      <tr key={ID.toString()}>
        <td>{ID}</td>
        <td>{name}</td>
        <td>{price}</td>
      </tr>
    );
  };

  return (
    <>
      <h1>Priceomatic Client</h1>
      <hr />
      <h3>Productos:</h3>
      <table className="table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nombre</th>
            <th>Precio</th>
          </tr>
        </thead>
        <tbody>{productos.map((product) => renderItem(product))}</tbody>
      </table>
    </>
  );
};
