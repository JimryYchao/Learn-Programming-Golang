<div id="img"><img src="../img/1.png" alt="Image description"></div>

<script>
    const imageContainer = document.getElementById('img');
    const menuItems = document.querySelectorAll('.menu li');

    menuItems.forEach(item => {
        item.addEventListener('mouseover', () => {
            imageContainer.style.display = 'block';
        });

        item.addEventListener('mouseout', () => {
            imageContainer.style.display = 'none';
        });
    });
</script>




## Package slice

Slice 包提供了一些对任意类型的切片的功能函数。

| Identifier                           | Description |
| :----------------------------------- | :---------- |
| <a id="myDIV"><img src="../img/1.png" alt="Image Description" title="Hover Text"><code>BinarySearchFunc</code></a> |              |

---
### 