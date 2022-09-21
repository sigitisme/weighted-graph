# Problem Statement
A driver is assigned to pick and deliver a food orders. But the driver have a difficulty to find the best way / route to deliver all of this orders. Your task is to find routes with the lowest expenses for the driver !

The first parameters is a 2D Array of Integer, orders assigned to the driver. Each orders will have this format : `[pick,destination]`.
pick, means driver must go to this location first to pick the delivery.
After driver arrives on pick location, then he will go to the destination.
The order must be done sequentially.
The second parameters is a 2D Array of Integer consisting of routes. Each items will represents bank transfer information in this format : `[from,to,cost]`.
from and to indicates the location driver can access to.
Cost is the amount expenses needed to use this route.
