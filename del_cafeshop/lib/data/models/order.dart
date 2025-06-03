class Order {
  final int id;
  final int? userId; // Nullable karena API mungkin tidak menyediakan user_id
  final String status;
  final DateTime createdAt;
  final DateTime updatedAt;
  final String customerName;
  final List<Product> products;

  Order({
    required this.id,
    required this.userId,
    required this.status,
    required this.createdAt,
    required this.updatedAt,
    required this.customerName,
    required this.products,
  });

  factory Order.fromJson(Map<String, dynamic> json) {
    return Order(
      id: json['order_id'] as int,
      userId: json['user_id'] as int?, // Nullable, null jika tidak ada
      status: json['status'] as String,
      createdAt: DateTime.parse(json['created_at'] ?? DateTime.now().toIso8601String()),
      updatedAt: DateTime.parse(json['updated_at'] ?? DateTime.now().toIso8601String()),
      customerName: json['customer_name'] as String,
      products: (json['products'] as List<dynamic>)
          .map((product) => Product.fromJson(product))
          .toList(),
    );
  }
}

class Product {
  final int productId;
  final String productName;
  final int price;
  final int quantity;

  Product({
    required this.productId,
    required this.productName,
    required this.price,
    required this.quantity,
  });

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      productId: json['product_id'] as int,
      productName: json['product_name'] as String,
      price: json['price'] as int,
      quantity: json['quantity'] as int,
    );
  }
}