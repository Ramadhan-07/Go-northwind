// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package gen

import (
	"database/sql"
)

type Category struct {
	CategoryID   int16          `db:"category_id" json:"categoryId"`
	CategoryName string `db:"category_name" json:"categoryName"`
	Description  string `db:"description" json:"description"`
	Picture      []byte         `db:"picture" json:"picture"`
}

type Customer struct {
	CustomerID   string         `db:"customer_id" json:"customerId"`
	CompanyName  sql.NullString `db:"company_name" json:"companyName"`
	ContactName  sql.NullString `db:"contact_name" json:"contactName"`
	ContactTitle sql.NullString `db:"contact_title" json:"contactTitle"`
	Address      sql.NullString `db:"address" json:"address"`
	City         sql.NullString `db:"city" json:"city"`
	Region       sql.NullString `db:"region" json:"region"`
	PostalCode   sql.NullString `db:"postal_code" json:"postalCode"`
	Country      sql.NullString `db:"country" json:"country"`
	Phone        sql.NullString `db:"phone" json:"phone"`
	Fax          sql.NullString `db:"fax" json:"fax"`
}

type Employee struct {
	EmployeeID      int16          `db:"employee_id" json:"employeeId"`
	LastName        sql.NullString `db:"last_name" json:"lastName"`
	FirstName       sql.NullString `db:"first_name" json:"firstName"`
	Title           sql.NullString `db:"title" json:"title"`
	TitleOfCourtesy sql.NullString `db:"title_of_courtesy" json:"titleOfCourtesy"`
	BirthDate       sql.NullTime   `db:"birth_date" json:"birthDate"`
	HireDate        sql.NullTime   `db:"hire_date" json:"hireDate"`
	Address         sql.NullString `db:"address" json:"address"`
	City            sql.NullString `db:"city" json:"city"`
	Region          sql.NullString `db:"region" json:"region"`
	PostalCode      sql.NullString `db:"postal_code" json:"postalCode"`
	Country         sql.NullString `db:"country" json:"country"`
	HomePhone       sql.NullString `db:"home_phone" json:"homePhone"`
	Extension       sql.NullString `db:"extension" json:"extension"`
	Photo           []byte         `db:"photo" json:"photo"`
	Notes           sql.NullString `db:"notes" json:"notes"`
	ReportTo        sql.NullInt16  `db:"report_to" json:"reportTo"`
	PhotoPath       sql.NullString `db:"photo_path" json:"photoPath"`
}

type Order struct {
	OrderID        int16           `db:"order_id" json:"orderId"`
	OrderDate      sql.NullTime    `db:"order_date" json:"orderDate"`
	RequiredDate   sql.NullTime    `db:"required_date" json:"requiredDate"`
	ShippedDate    sql.NullTime    `db:"shipped_date" json:"shippedDate"`
	Freight        sql.NullFloat64 `db:"freight" json:"freight"`
	ShipName       sql.NullString  `db:"ship_name" json:"shipName"`
	ShipAddress    sql.NullString  `db:"ship_address" json:"shipAddress"`
	ShipCity       sql.NullString  `db:"ship_city" json:"shipCity"`
	ShipRegion     sql.NullString  `db:"ship_region" json:"shipRegion"`
	ShipPostalCode sql.NullString  `db:"ship_postal_code" json:"shipPostalCode"`
	ShipCountry    sql.NullString  `db:"ship_country" json:"shipCountry"`
	EmployeeID     sql.NullInt16   `db:"employee_id" json:"employeeId"`
	CustomerID     sql.NullString  `db:"customer_id" json:"customerId"`
	ShipVia        sql.NullInt16   `db:"ship_via" json:"shipVia"`
}

type OrderDetail struct {
	OrderID   int16           `db:"order_id" json:"orderId"`
	ProductID int16           `db:"product_id" json:"productId"`
	UnitPrice sql.NullFloat64 `db:"unit_price" json:"unitPrice"`
	Quantity  sql.NullInt16   `db:"quantity" json:"quantity"`
	Discount  sql.NullFloat64 `db:"discount" json:"discount"`
}

type Product struct {
	ProductID       int16           `db:"product_id" json:"productId"`
	ProductName     sql.NullString  `db:"product_name" json:"productName"`
	QuantityPerUnit sql.NullString  `db:"quantity_per_unit" json:"quantityPerUnit"`
	UnitPrice       sql.NullFloat64 `db:"unit_price" json:"unitPrice"`
	UnitsInStock    sql.NullInt16   `db:"units_in_stock" json:"unitsInStock"`
	UnitsInOrder    sql.NullInt16   `db:"units_in_order" json:"unitsInOrder"`
	ReorderLevel    sql.NullInt16   `db:"reorder_level" json:"reorderLevel"`
	Discontinued    sql.NullInt32   `db:"discontinued" json:"discontinued"`
	SupplierID      sql.NullInt16   `db:"supplier_id" json:"supplierId"`
	CategoryID      sql.NullInt16   `db:"category_id" json:"categoryId"`
}

type Shipper struct {
	ShipperID   int16          `db:"shipper_id" json:"shipperId"`
	CompanyName sql.NullString `db:"company_name" json:"companyName"`
	Phone       sql.NullString `db:"phone" json:"phone"`
}

type Supplier struct {
	SupplierID   int16          `db:"supplier_id" json:"supplierId"`
	CompanyName  sql.NullString `db:"company_name" json:"companyName"`
	ContactName  sql.NullString `db:"contact_name" json:"contactName"`
	ContactTitle sql.NullString `db:"contact_title" json:"contactTitle"`
	Address      sql.NullString `db:"address" json:"address"`
	City         sql.NullString `db:"city" json:"city"`
	Region       sql.NullString `db:"region" json:"region"`
	PostalCode   sql.NullString `db:"postal_code" json:"postalCode"`
	Country      sql.NullString `db:"country" json:"country"`
	Phone        sql.NullString `db:"phone" json:"phone"`
	Fax          sql.NullString `db:"fax" json:"fax"`
	Homepage     sql.NullString `db:"homepage" json:"homepage"`
}
