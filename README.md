# polygons

## Backface Culling

**N**: surface normal perpendicular to the surface  
**V**: view vector from the surface to the viewer  

If a surface is in our view, **N** · **V** >= 0.  

If P0, P1, P2 are the three points of a triangle in counterclockwise order,  
**A** = <P1 - P0>  
**B** = <P2 - P0>  
**N** = **A** x **B**
