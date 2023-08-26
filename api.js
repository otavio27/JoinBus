function distance1(lat1, lon1, lat2, lon2, unit) {
    const radlat1 = Math.PI * lat1 / 180;
    const radlat2 = Math.PI * lat2 / 180;
    const radlon1 = Math.PI * lon1 / 180;
    const radlon2 = Math.PI * lon2 / 180;
    const theta = lon1 - lon2;
    const radtheta = Math.PI * theta / 180;
    let dist = Math.sin(radlat1) * Math.sin(radlat2) + Math.cos(radlat1) * Math.cos(radlat2) * Math.cos(radtheta);
    dist = Math.acos(dist);
    dist = dist * 180 / Math.PI;
    dist = dist * 60 * 1.1515;
    if (unit === "K") { dist = dist * 1.609344 }
    if (unit === "N") { dist = dist * 0.8684 }
    return dist.toFixed(2);
}

//const distanceInKm = distance(-26.3250729682203, -48.7949848, -26.3183498, -48.7899487, "K");


function distance2(lat1, lon1, lat2, lon2, unit) {
    const radlat1 = Math.PI * lat1 / 180;
    const radlat2 = Math.PI * lat2 / 180;
    const theta = lon1 - lon2;
    const radtheta = Math.PI * theta / 180;
    let dist = Math.sin(radlat1) * Math.sin(radlat2) + Math.cos(radlat1) * Math.cos(radlat2) * Math.cos(radtheta);
    dist = Math.acos(dist);
    dist = dist * 180 / Math.PI;
    dist = dist * 60 * 1.1515;
    if (unit === "K") { dist = dist * 1.609344 }
    if (unit === "N") { dist = dist * 0.8684 }
    return dist.toFixed(2);
}

function calculateDistance(lat1, long1, lat2, long2) {

    //radians
    lat1 = (lat1 * 2.0 * Math.PI) / 60.0 / 360.0;
    long1 = (long1 * 2.0 * Math.PI) / 60.0 / 360.0;
    lat2 = (lat2 * 2.0 * Math.PI) / 60.0 / 360.0;
    long2 = (long2 * 2.0 * Math.PI) / 60.0 / 360.0;


    // use to different earth axis length    
    var a = 6378137.0;        // Earth Major Axis (WGS84)    
    var b = 6356752.3142;     // Minor Axis    
    var f = (a - b) / a;        // "Flattening"    
    var e = 2.0 * f - f * f;      // "Eccentricity"      

    var beta = (a / Math.sqrt(1.0 - e * Math.sin(lat1) * Math.sin(lat1)));
    var cos = Math.cos(lat1);
    var x = beta * cos * Math.cos(long1);
    var y = beta * cos * Math.sin(long1);
    var z = beta * (1 - e) * Math.sin(lat1);

    beta = (a / Math.sqrt(1.0 - e * Math.sin(lat2) * Math.sin(lat2)));
    cos = Math.cos(lat2);
    x -= (beta * cos * Math.cos(long2));
    y -= (beta * cos * Math.sin(long2));
    z -= (beta * (1 - e) * Math.sin(lat2));

    return (Math.sqrt((x * x) + (y * y) + (z * z)) / 1000);
}

const distanceInKm1 = distance1(-26.3250729, -48.7949848, -26.3183498, -48.7899487, "K");
console.log(`A distância 1 entre ônibus e você é de ${distanceInKm1} km.`);

const distanceInKm2 = distance1(-26.3250729, -48.7949848, -26.3183498, -48.7899487, "K");
console.log(`A distância 2 entre ônibus e você é de ${distanceInKm2} km.`);

const distanceInKm = calculateDistance(-26.3250729, -48.7949848, -26.3183498, -48.7899487);

console.log(`A distância é de ${distanceInKm.toFixed(2)} km`);